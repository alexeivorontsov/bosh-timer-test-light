This assumes that you have bosh Director deployed locally and it is managing a local bosh installation.

##Deployment

Step 0. Make a local working directory and change it to be you current working dir

    mkdir bosh-test
    cd bosh-test
    
Step 1. Clone this repository

    git clone https://github.com/alexeivorontsov/bosh-timer-test-light.git
    
Step 2. Point bosh director to the cloned deployment manifest file

    cd bosh-timer-test-light
    ls -la
    bosh deployment manifest.yml
    bosh status

Step 3. Download and save locally a distribution of golang (from here https://storage.googleapis.com/golang/go1.5.2.linux-amd64.tar.gz). For example to /tmp/golang-blobs/go1.5.2.linux-amd64.tar.gz file and add this blob to bosh

    bosh add blob /tmp/golang-blobs/go1.5.2.linux-amd64.tar.gz golang
    bosh blobs
    
Step 4. Create a new release

    bosh create release --force


Step 5. Upload the new release to bosh (accept default release name when prompted)

    bosh upload release    
    bosh releases
    
Step 6. Deploy the new release to bosh (answer 'yes' when prompted)

    bosh deploy
    bosh deployments
    
Step 7. Check the log file for the running application. Login into the VM running the test timer and check the log files. Exit the VM when finished

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

