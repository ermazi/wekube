clean:
	@echo "clean"
	@find ./ -size +3M|xargs rm -f
	@echo "clean done"