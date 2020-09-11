ETCD_VERSION=v3.2.18
TOKEN=my-etcd-token
CLUSTER_STATE=new
NAME_1=etcd-node-0
HOST_1=10.9.56.153
CLUSTER=${NAME_1}=http://${HOST_1}:23801

# 对于节点1
THIS_NAME=${NAME_1}
THIS_IP=${HOST_1}
sudo docker run -d --net=host --name ${THIS_NAME} dockerhub.qingcloud.com/coreos/etcd:${ETCD_VERSION} \
    /usr/local/bin/etcd \
    --data-dir=data.etcd --name ${THIS_NAME} \
    --initial-advertise-peer-urls http://${THIS_IP}:23801 --listen-peer-urls http://${THIS_IP}:23801 \
    --advertise-client-urls http://${THIS_IP}:23791 --listen-client-urls http://${THIS_IP}:23791 \
    --initial-cluster ${CLUSTER} \
    --initial-cluster-state ${CLUSTER_STATE} --initial-cluster-token ${TOKEN}
