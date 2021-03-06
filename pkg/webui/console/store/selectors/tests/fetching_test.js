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

import { createFetchingSelector } from '@ttn-lw/lib/store/selectors/fetching'

describe('Fetching selectors', () => {
  const BASE_ACTION_TYPE = 'BASE_ACTION'
  let initialState = null

  describe('when created with a single base action type', () => {
    const selector = createFetchingSelector(BASE_ACTION_TYPE)

    beforeAll(() => {
      initialState = { ui: { fetching: {} } }
    })

    describe('when it has no fetching entries', () => {
      it('returns `false`', () => {
        expect(selector(initialState)).toBe(false)
      })
    })

    describe('when it has fetching entry', () => {
      beforeAll(() => {
        initialState.ui.fetching[BASE_ACTION_TYPE] = true
      })

      it('return `true`', () => {
        expect(selector(initialState)).toBe(true)
      })
    })
  })

  describe('when created with two base action types', () => {
    const BASE_ACTION_TYPE_OTHER = 'BASE_ACTION_OTHER'
    const selector = createFetchingSelector([BASE_ACTION_TYPE, BASE_ACTION_TYPE_OTHER])

    beforeAll(() => {
      initialState = { ui: { fetching: {} } }
    })

    describe('when it has no fetching entries', () => {
      it('return `false`', () => {
        expect(selector(initialState)).toBe(false)
      })
    })

    describe('when it has a fetching entry', () => {
      beforeAll(() => {
        initialState.ui.fetching[BASE_ACTION_TYPE] = true
      })

      it('return `true`', () => {
        expect(selector(initialState)).toBe(true)
      })

      describe('when it has two fetching entries', () => {
        beforeAll(() => {
          initialState.ui.fetching[BASE_ACTION_TYPE_OTHER] = true
        })

        it('return `true`', () => {
          expect(selector(initialState)).toBe(true)
        })
      })
    })
  })
})
