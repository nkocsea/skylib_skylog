wrun:
	go run src/main.go

run:
	reflex -r '\.go' -s -- sh -c "go run src/main.go"

tags:
	git ls-remote --tags

commit:
	git status
	git add .
	git commit -m"$m"
	git push
	git tag $t
	git push --tags