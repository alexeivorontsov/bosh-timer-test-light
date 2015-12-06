This assumes that you have bosh Director deployed locally and it is managing a local bosh installation. These instructions were written and tested for Mac OS X.

##Deployment

Step 1. Make a local working directory and change it to be you current working dir

    mkdir bosh-test
    cd bosh-test
    
Step 2. Clone this repository

    git clone https://github.com/alexeivorontsov/bosh-timer-test-light.git
    
Step 3. Set you bosh Director UUID in the deployment manifest file. You can do that manually with a text editor or you can run this command to accomplish that

    sed -i '' "s|REPLACE_THIS_TEXT_WITH_YOUR_DIRECTOR_UUID|$(bosh status --uuid)|" manifest.yml

Open the deployment manifest file (manifest.yml) and change the following line (line 3)

    director_uuid: 238e1b09-f66d-4432-9101-fd0d03efaed3
    
Step 4. Point bosh director to the cloned deployment manifest file

    cd bosh-timer-test-light
    ls -la
    bosh deployment manifest.yml
    bosh status


Step 5. Get and upload the stemcell that is used for deployment. It may already be installed on your local bosh director

    bosh upload stemcell https://bosh.io/d/stemcells/bosh-warden-boshlite-ubuntu-trusty-go_agent --skip-if-exists
    bosh stemcells
    
Step 6. Download and save locally a distribution of golang (from here https://storage.googleapis.com/golang/go1.5.2.linux-amd64.tar.gz). For example to /tmp/golang-blobs/go1.5.2.linux-amd64.tar.gz file and add this blob to bosh

    bosh add blob /tmp/golang-blobs/go1.5.2.linux-amd64.tar.gz golang
    bosh blobs
    
Step 7. Create a new release

    bosh create release --force


Step 8. Upload the new release to bosh (accept default release name when prompted)

    bosh upload release    
    bosh releases
    
Step 9. Deploy the new release to bosh (answer 'yes' when prompted)

    bosh deploy
    bosh deployments
    
Step 10. Check the log file for the running application. Login into the VM running the test timer and check the log files. Exit the VM when finished

    bosh ssh timer_runner/0
    cd /var/vcap/sys/log/timer_runner/
    tail -f server.stdout.log
    
## Clean up

Remove the new deployment from bosh (answer 'yes' when prompted)

    bosh delete deployment my-new-timer-deployment --force
    
Remove the new release from bosh (answer 'yes' when prompted)

    bosh delete release my-new-timer-release --force
    
Verify that both the deployment and the release are deleted

    bosh deployments
    bosh releases
    
Remove the working directory
  
    cd ..
    rm -rf ./bosh-timer-test-light

