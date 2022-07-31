/**
 * @jest-environment jsdom
 */
import { mount } from '@vue/test-utils'

import HelloWorld from './HelloWorld.vue'

describe('HelloWorld.vue', () => {
  it('renders Test text', () => {
    const msg = 'Test text'
    const wrapper = mount(HelloWorld, {
      props: { msg },
    })

    expect(wrapper.text()).toContain(msg)
  })
})
