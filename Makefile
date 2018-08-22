clean:
	@echo "clean all executable files"
	@find ./ -size +3M|xargs rm -f
	@echo "clean done"