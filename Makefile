TARGET = feishu_bot
LOG_DIR = ./.log/
OUTPUT_DIR = ./output

.PHONY: build
build:
	go build -o $(OUTPUT_DIR)/$(TARGET)

.PHONY: run
run: build
	$(OUTPUT_DIR)/$(TARGET) --env=dev

.PHONY: realse
realse: build
	mkdir -p $(LOG_DIR)
	nohup $(OUTPUT_DIR)/$(TARGET) --env=pro > $(LOG_DIR)/$(TARGET).log 2>&1 &

.PHONY: show
show:
	ps -ef | grep $(TARGET)

.PHONY: log
log:
	tail $(LOG_DIR)/$(TARGET).log

process_id = $(shell ps -ef | grep $(TARGET) | awk 'NR==1{print $$2}')

.PHONY: kill
kill:
	sudo -S kill ${process_id}