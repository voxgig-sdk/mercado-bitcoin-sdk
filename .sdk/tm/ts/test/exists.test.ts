
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { MercadoBitcoinSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await MercadoBitcoinSDK.test()
    equal(null !== testsdk, true)
  })

})
