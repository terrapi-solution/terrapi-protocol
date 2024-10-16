#
# Licensed to the TerrAPI under one or more contributor
# license agreements. See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership. The TerrAPI licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

PROTO_NAMES=(
    "activity"
    "deployment"
    "health"
)

for name in "${PROTO_NAMES[@]}"; do
  protoc ${name}/v1/${name}.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=.
  if [ $? -ne 0 ]; then
      echo "error processing ${name}.proto"
      exit $?
  fi
done