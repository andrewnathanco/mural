# Configure Docker to use the Artifact Registry
gcloud auth configure-docker us-east1-docker.pkg.dev

# build the image
docker build -t us-east1-docker.pkg.dev/ancgames-prod/ancgames-docker/ancgames-www:latest --build-arg SUPABASE_URL=${SUPABASE_URL} --build-arg SUPABASE_KEY=${SUPABASE_KEY} .

