package kubescaler

const userDataTpl = `#cloud-config

hostname: ""
ssh_authorized_keys:
  - "{{ .SSHPubKey }}"
write_files:
  - path: "/opt/bin/download-k8s-binary"
    permissions: "0755"
    content: |
      #!/bin/bash
      source /etc/environment
      K8S_VERSION={{ .KubeVersion }}
      mkdir -p /opt/bin
      mkdir /etc/multipath/
      touch /etc/multipath/bindings
      curl -sSL -o /opt/bin/kubectl https://storage.googleapis.com/kubernetes-release/release/${K8S_VERSION}/bin/linux/amd64/kubectl
      chmod +x /opt/bin/$FILE
      chmod +x /opt/bin/kubectl

      curl -sSL -o /opt/bin/cni.tar.gz https://storage.googleapis.com/kubernetes-release/network-plugins/cni-07a8a28637e97b22eb8dfe710eeae1344f69d16e.tar.gz
      tar xzf "/opt/bin/cni.tar.gz" -C "/opt/bin" --overwrite
      mv /opt/bin/bin/* /opt/bin
      rm -r /opt/bin/bin/
      rm -f "/opt/bin/cni.tar.gz"

      cd /opt/bin/
      git clone https://github.com/packethost/packet-block-storage.git
      cd packet-block-storage
      chmod 755 ./*
      /opt/bin/packet-block-storage/packet-block-storage-attach

      cd /tmp
      wget https://github.com/digitalocean/doctl/releases/download/v1.4.0/doctl-1.4.0-linux-amd64.tar.gz
      tar xf /tmp/doctl-1.4.0-linux-amd64.tar.gz
      sudo mv /tmp/doctl /opt/bin/
      sudo mkdir -p /root/.config/doctl/
      sudo touch /root/.config/doctl/config.yaml
  - path: "/etc/kubernetes/manifests/kube-proxy.yaml"
    permissions: "0644"
    owner: "root"
    content: |
      apiVersion: v1
      kind: Pod
      metadata:
        name: kube-proxy
        namespace: kube-system
      spec:
        hostNetwork: true
        containers:
        - name: kube-proxy
          image: gcr.io/google_containers/hyperkube:{{ .KubeVersion }}
          command:
          - /hyperkube
          - proxy
          - --master=https://{{ .MasterPrivateAddr }}:{{ .KubeAPIPort }}
          - --kubeconfig=/etc/kubernetes/worker-kubeconfig.yaml
          - --v=2
          - --proxy-mode=iptables
          securityContext:
            privileged: true
          volumeMounts:
            - mountPath: /etc/ssl/certs
              name: "ssl-certs"
            - mountPath: /etc/kubernetes/worker-kubeconfig.yaml
              name: "kubeconfig"
              readOnly: true
            - mountPath: /etc/kubernetes/ssl
              name: "etc-kube-ssl"
              readOnly: true
        volumes:
          - name: "ssl-certs"
            hostPath:
              path: "/usr/share/ca-certificates"
          - name: "kubeconfig"
            hostPath:
              path: "/etc/kubernetes/worker-kubeconfig.yaml"
          - name: "etc-kube-ssl"
            hostPath:
              path: "/etc/kubernetes/ssl"
  - path: "/etc/kubernetes/worker-kubeconfig.yaml"
    permissions: "0644"
    owner: "root"
    content: |
      apiVersion: v1
      kind: Config
      users:
      - name: kubelet
        user:
          token: {{ .KubeAPIPassword }}
      clusters:
      - name: local
        cluster:
          insecure-skip-tls-verify: true
          server: https://{{ .MasterPrivateAddr }}:{{ .KubeAPIPort }}
      contexts:
      - context:
          cluster: local
          user: kubelet
        name: service-account-context
      current-context: service-account-context
coreos:
  update:
    reboot-strategy: off
  flannel:
    iface: $COREOS_PRIVATE_IPV4
    etcd_endpoints: http://{{ .MasterPrivateAddr }}:2379
  units:
    - name: "flanneld.service"
      command: start
      drop-ins:
        - name: 50-network-config.conf
          content: |
            [Unit]
            Requires=etcd-member.service
            [Service]
            Environment=FLANNEL_IMAGE_TAG=v0.9.0
            ExecStartPre=/usr/bin/etcdctl set /coreos.com/network/config '{"Network":"10.2.0.0/16", "Backend": {"Type": "vxlan"}}' 
    - name: "docker.service"
      command: start
      drop-ins:
        - name: 40-flannel.conf
          content: |
            [Unit]
            Requires=flanneld.service
            After=flanneld.service
    - name: iscsid.service
      enable: true
      command: start
    - name: kubelet.service
      command: start
      content: |
        [Unit]
        Description=Kubernetes Kubelet Server
        Documentation=https://github.com/kubernetes/kubernetes
        Requires=docker.service network-online.target
        After=docker.service network-online.target

        [Service]
        ExecStartPre=/bin/mkdir -p /var/lib/kubelet
        ExecStartPre=/bin/mount --bind /var/lib/kubelet /var/lib/kubelet
        ExecStartPre=/bin/mount --make-shared /var/lib/kubelet
        ExecStart=/usr/bin/docker run \
                --net=host \
                --pid=host \
                --privileged \
                -v /dev:/dev \
                -v /sys:/sys:ro \
                -v /var/run:/var/run:rw \
                -v /var/lib/docker/:/var/lib/docker:rw \
                -v /var/lib/kubelet/:/var/lib/kubelet:shared \
                -v /var/log:/var/log:shared \
                -v /srv/kubernetes:/srv/kubernetes:ro \
                -v /etc/kubernetes:/etc/kubernetes:ro \
                gcr.io/google-containers/hyperkube:{{ .KubeVersion }} \
                /hyperkube kubelet --allow-privileged=true \
                --cluster-dns=10.3.0.10 \
                --cluster_domain=cluster.local \
                --pod-manifest-path=/etc/kubernetes/manifests \
                --kubeconfig=/etc/kubernetes/worker-kubeconfig.yaml \
                --volume-plugin-dir=/etc/kubernetes/volumeplugins \
                --cloud-provider={{ .ProviderName }} \
                --register-node=true
        Restart=always
        StartLimitInterval=0
        RestartSec=10
        KillMode=process

        [Install]
        WantedBy=multi-user.target
`
