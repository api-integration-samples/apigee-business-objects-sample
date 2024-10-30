SERVICE_NAME=business-objects-service

# deploy to cloud run
gcloud run deploy $SERVICE_NAME --source . --project $PROJECT_ID --region $REGION --allow-unauthenticated

# get service url
SERVICE_URL=$(gcloud run services describe $SERVICE_NAME --format 'value(status.url)' --region $REGION --project $PROJECT_ID)
