output_dir = output
artifact = $(output_dir)/kubecfg

build : clean output test
	go build -o $(artifact)

clean :
	rm -rf $(output_dir)

output :
	mkdir $(output_dir)

test :
	go test ./...

bash_completion_dir = ~/.bash_completion.d
bash_completion_target = $(bash_completion_dir)/kubecfg.bash_completion
bash_completion_source = scripts/kubecfg.bash_completion

# yes | cp -rf -> force overwrite
install : build
	yes | sudo cp -rf $(artifact) /usr/local/bin
	yes | cp -rf $(bash_completion_source) $(bash_completion_target)
