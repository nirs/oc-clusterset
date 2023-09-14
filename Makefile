# SPDX-FileCopyrightText: The RamenDR authors
# SPDX-License-Identifier: Apache-2.0

output := oc-clusterset

all:
	go build -o $(output)

test: reuse lint

lint:
	golangci-lint run

reuse:
	reuse lint

clean:
	rm -f $(output)
