<template>
  <div v-if="dropdown.length" class="logo-dropdown" id="logo-dropdown" :class="{ open: isOpen }" @click.stop>
    <div class="logo-dd-top">
      <a v-for="link in dropdown[0].links" :key="link.text" :href="link.href">{{ link.text }}</a>
    </div>
    <div v-for="group in dropdown.slice(1)" :key="group.section" class="divider">
      <div class="logo-dd-section">{{ group.section }}</div>
      <div class="logo-dd-grid">
        <a v-for="link in group.links" :key="link.text" :href="link.href">
          {{ link.text }}
          <span v-if="link.badge">{{ link.badge }}</span>
        </a>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useFetch } from '../composables/useFetch.js'

defineProps(['isOpen'])

const { data: dropdown } = useFetch('/api/dropdown', [])
</script>