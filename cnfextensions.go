// Copyright 2022 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cnfextensions

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/sirupsen/logrus"
	"github.com/test-network-function/cnf-certification-test/cnf-certification-test/identifiers"
	"github.com/test-network-function/cnf-certification-test/cnf-certification-test/results"
	"github.com/test-network-function/cnf-certification-test/pkg/provider"
	"github.com/test-network-function/cnf-certification-test/pkg/testhelper"
	"github.com/test-network-function/test-network-function-claim/pkg/claim"
)

const (
	ExampleTestSuite = "cnfextensions"
)

// Simple example test suite. More complex scenarios should use subdirectories
var _ = ginkgo.Describe(ExampleTestSuite, func() {
	_, TestExampleIdentifier := claim.BuildTestCaseDescription("example", ExampleTestSuite, `this is an example test case`, `this is what todo to pass the test case`, `normative`, "No Documented Process", "http://myorg.org/testplan Section 5.2", false, "mytag", "pre_release")
	var env provider.TestEnvironment
	ginkgo.BeforeEach(func() {
		env = provider.GetTestEnvironment()
	})
	ginkgo.ReportAfterEach(results.RecordResult)

	testID, tags := identifiers.GetGinkgoTestIDAndLabels(TestExampleIdentifier)
	ginkgo.It(testID, ginkgo.Label(tags...), func() {
		testhelper.SkipIfEmptyAny(ginkgo.Skip, env.Pods)
		testExample(&env)
	})
})

func testExample(env *provider.TestEnvironment) {
	logrus.Infof("Hello! There are %d pods under test", len(env.Pods))
}
