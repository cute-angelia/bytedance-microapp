.PHONY: up tag

up:
	git pull origin main
	git add .
	git commit -am "update"
	git push origin main
	@echo "\n 发布中..."

tag:
	git pull origin master
	git add .
	git commit -am "upload"
	git push origin main
	git tag v1.0.0
	git push --tags
	@echo "\n tags 发布中..."
