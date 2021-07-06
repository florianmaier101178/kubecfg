app_name = kubecfg
output_dir = output
artifact = $(output_dir)/$(app_name)

build : clean output test
	go build -o $(artifact)

clean :
	rm -rf $(output_dir)

output :
	mkdir $(output_dir)

test :
	go test ./...

bash_completion_dir = ~/.bash_completion.d
bash_completion_target = $(bash_completion_dir)/$(app_name).bash_completion
bash_completion_source = scripts/$(app_name).bash_completion

# yes | cp -rf -> force overwrite
install : build
	yes | sudo cp -f $(artifact) /usr/local/bin
	yes | cp -f $(bash_completion_source) $(bash_completion_target)

uninstall :
	yes | sudo rm -f /usr/local/bin/$(app_name)
	yes | rm -f $(bash_completion_target)
