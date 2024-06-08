cover:
	go test -cover ./...

test:
	go test ./...


code-analysis:
	@cd ..; echo "Обратите внимание на следующие файлы, возможно им требуется рефакторинг: " \
 	&& echo "Цикломатический анализ" && gocyclo -over 20 -ignore "_test|_mock" . \
 	&& echo "Когнтивный анализ" && gocognit -over 10 -ignore "_test|_mock" .

business-logic-analysis:
	@cd ..; echo "Самые сложные функции в уровне бизнесс-логики" \
 	&& echo "Цикломатический анализ" && gocyclo  -over 20 -ignore "_test|_mock" ./cart/internal/service \
	&& echo "Когнитивный анализ" && gocognit -top 10 -ignore "_test|_mock" ./cart/internal/service

