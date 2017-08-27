// Copyright 2017 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package casbin

// GetRolesForUserInDomain gets the roles that a user has inside a domain.
func (e *Enforcer) GetRolesForUserInDomain(name string, domain string) []string {
	return e.model["g"]["g"].RM.GetRoles(name, domain)
}

// GetPermissionsForUserInDomain gets permissions for a user or role inside a domain.
func (e *Enforcer) GetPermissionsForUserInDomain(user string, domain string) [][]string {
	return e.GetFilteredPolicy(0, user, domain)
}