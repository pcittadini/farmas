image:	
		cp Dockerfile ./skydns
		cd skydns; docker build --no-cache -t skydns .
		docker tag -f skydns snowpunk/skydns:$(v)
		docker push snowpunk/skydns:$(v)
