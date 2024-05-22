echo "* (SCRIPT) Triton agent start."
echo "* (SCRIPT) Triton agent docker build."
docker build -t triton-agent .
if [ $? -ne 0 ]; then
	echo "*** (ERROR) Is Docker running?"
	exit 1
fi
echo "* (SCRIPT) Docker build success!"

echo "* (SCRIPT) Triton agent docker start."
docker run -it --rm --name triton-agent --network triton -p 7000:7000 -v $PWD/../2/triton-server/models:/models triton-agent
echo "* (SCRIPT) Exit."