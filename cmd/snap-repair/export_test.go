// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2017 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package main

import (
	"gopkg.in/retry.v1"
)

var (
	Parser    = parser
	ParseArgs = parseArgs
	Run       = run
)

func MockFetchRetryStrategy(strategy retry.Strategy) (restore func()) {
	originalFetchRetryStrategy := fetchRetryStrategy
	fetchRetryStrategy = strategy
	return func() {
		fetchRetryStrategy = originalFetchRetryStrategy
	}
}

func MockPeekRetryStrategy(strategy retry.Strategy) (restore func()) {
	originalPeekRetryStrategy := peekRetryStrategy
	peekRetryStrategy = strategy
	return func() {
		peekRetryStrategy = originalPeekRetryStrategy
	}
}

func MockMaxRepairScriptSize(maxSize int) (restore func()) {
	originalMaxSize := maxRepairScriptSize
	maxRepairScriptSize = maxSize
	return func() {
		maxRepairScriptSize = originalMaxSize
	}
}
