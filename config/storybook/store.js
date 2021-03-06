// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { createStore, applyMiddleware, compose, combineReducers } from 'redux'
import { routerMiddleware, connectRouter } from 'connected-react-router'

const createRootReducer = history =>
  combineReducers({
    router: connectRouter(history),
  })

export default history => {
  const middleware = applyMiddleware(routerMiddleware(history))

  return createStore(createRootReducer(history), compose(middleware))
}
