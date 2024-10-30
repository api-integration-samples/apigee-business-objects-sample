# Apigee Business Objects Sample
This is a sample project of business objects (similar to SAP data in REST format), deployed to a Cloud Run service, and offered through Apigee APIs.

## Deployment
Deployment is very simple with the **gcloud** CLI installed.

```sh
# first copy and set environment variables
cp 1.env.sh 1.env.local.sh
# edit and save local env variables
nano 1.env.local.sh

# now deploy to cloud run
./2.deploy-cloudrun.sh

# set as server in local oas file
cp oas.orderservice.yaml oas.orderservice.local.yaml
sed -i "/  - url: /c\  - url: $SERVICE_URL" oas.orderservice.local.yaml

# deploy apigee proxy based on local oas file
apigeecli apis create openapi -o $PROJECT_ID -e $APIGEE_ENV --ovr -n SalesOrderApi-v1 -p /orderservice --oas-base-folderpath . --oas-name oas.orderservice.local.yaml -t $(gcloud auth print-access-token)

# get apigee hostname for environment
HOST_NAME=$(apigeecli envgroups list -o $PROJECT_ID -t $(gcloud auth print-access-token) | jq --raw-output '.environmentGroups[] | select(.name == '\""$APIGEE_ENV\""') | .hostnames[0]')

# call API after a few seconds
curl "$HOST_NAME/orderservice/orders"
# should return order data {"orders": [...]}

```