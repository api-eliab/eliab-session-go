#! /bin/bash

# Run: ./docker_build_image.sh 0.1.1

version="$1"
imagename=$(basename "$PWD")
localPath="../eliab-docker-img"
server="root@159.203.93.24"
remotePath="/root/docker-images"
appRemotePath="/root/api-eliab-dev"

echo -e "\n###### DEPLOY MODULE ######\n"

currentDate=`date +"%b-%d.%H.%M"`

echo "Build docker image in" $currentDate

filename=$imagename-v$version-$currentDate".tar"
echo "File name: "$filename

echo "Generate local docker image:"
docker build -t $imagename:$version . -f ./Dockerfile --rm=true || exit 1

docker rmi -f $(docker images -f dangling=true -q)

echo "Save docker image in file *.tar:"
docker save "$imagename":"$version" > $localPath/$filename || exit 1
echo $filename

echo "Upload docker image to server:"
rsync -rPavzh $localPath/$filename $server:$remotePath/$filename || exit 1
echo "###### SUCCESS LOAD ######"

loadImage="docker load < "$remotePath"/"$filename

echo "Enter server: "
ssh $server<< EOF
    $loadImage
    cd $appRemotePath
    docker-compose down
    docker-compose up &> eliabc.log&
    docker image prune -f
EOF
echo "\n###### SUCCESS DEPLOY MODULE ######\n"

git status

echo "Want you commit changes? [y/n]"
read commit

if [ $commit == 'y' ]
then 
    echo "Commit message: "
    read message
    git add .
    git commit -m "$message"
    git push
    echo "\n###### SUCCESS SCRIPT ######\n"
else 
    echo "\n###### SUCCESS SCRIPT[NOT COMMIT] ######\n"
    exit 1
fi