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

import reducer from '../ui/fetching'

describe('Fetching reducers', () => {
  const BASE = 'BASE_ACTION'
  const REQUEST = `${BASE}_REQUEST`
  const SUCCESS = `${BASE}_SUCCESS`
  const FAILURE = `${BASE}_FAILURE`

  it('return the initial state', () => {
    expect(reducer(undefined, {})).toEqual({})
  })

  describe('when dispatching invalid action type', () => {
    it('ignore the action', () => {
      expect(reducer({}, { type: BASE })).toEqual({})
    })
  })

  describe('when dispatching the `request` action', () => {
    it('set fetching to `true`', () => {
      expect(reducer({}, { type: REQUEST })).toEqual({
        [BASE]: true,
      })
    })

    describe('when dispatching the `success` action', () => {
      it('set fetching to `false`', () => {
        expect(reducer({}, { type: SUCCESS })).toEqual({
          [BASE]: false,
        })
      })
    })

    describe('when dispatching the `failure` action', () => {
      it('set fetching to `false`', () => {
        expect(reducer({}, { type: FAILURE })).toEqual({
          [BASE]: false,
        })
      })
    })
  })
})
