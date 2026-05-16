<template>
  <div class="footer" v-if="visible" v-show="!loading">
    <div class="footer-info">
      <div class="footer-content">
        <p>{{ footerData.text }}</p>
        <button class="footer-btn-ads join-nomads" @click="$emit('openPopup')">
          {{ footerData.btnText }}
        </button>
      </div>
      <button class="close-btn" @click="visible = false">✕</button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useFetch } from '../composables/useFetch.js'

defineEmits(['openPopup'])

const visible = ref(true)
const { data: footerData, loading } = useFetch('/api/footer', { text: '', btnText: '' })
</script>
