<template>
  <div
    class = '
      d-flex
    '
  >
    <v-text-field
      :disabled = '!getPrivateKey'
      dense
      v-model = 'shellCode'
      outlined
      label = 'Enter a command...'
      color = 'rgba(189, 147, 249, 1)'
      @change = 'issueMessages'
    ></v-text-field>

    <v-btn
      outlined
      color = 'rgba(189, 147, 249, 1)'
      class = '
        ml-4
      '
      @click = 'checkForReplies'
    >
      Check For Replies
    </v-btn>

    <h3
      v-if = '!getPrivateKey'
    >
      Please, go to Settings -> Identity and enter your private key in order to sign the messages.
    </h3>
  </div>
</template>

<script lang = 'ts'>
import Vue from 'vue'

import * as Service from '@/src/services'

import { mapGetters } from 'vuex'
import EventBus from '@/src/EventBus'

export default Vue.extend({
  data: () => ({
    shellCode: '',
    encryptPayload: false,
  }),

  props: {
    clients: {
      required: true,
    }
  },

  computed: {
    ...mapGetters([
      'getPrivateKey',
      'getAtilaHost',
      'getAtilaToken',
    ])
  },

  methods: {
    checkForReplies () {
      EventBus.$emit('interactionDialogUpdateSelectedClient', this.clients[0].id)
    },

    async issueMessages () {
      const message = await Service.Crypto.sign(this.getPrivateKey, { shell: this.shellCode })

      await new Service.Atila(this.$axios, this.getAtilaHost, this.getAtilaToken).publishMessage(this.clients[0].id, 'instruction', message)
      // const createdMessages = await Promise.all(this.clients.map(client => {
      //   return Service.Atila.publishMessage(this.$axios, client.id, 'instruction', message)
      // }))

      EventBus.$emit('interactionDialogUpdateSelectedClient', this.clients[0].id)

      this.shellCode = ''
    }
  },
})
</script>

<style scoped>
</style>