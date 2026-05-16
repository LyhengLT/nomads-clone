<template>
  <div class="news-quotes" v-show="!loading">
    <div v-for="item in hero.quotes" :key="item.alt"
         :class="['news-quote', item.extraClass]">
      <p>{{ item.quote }}</p>
      <img :src="item.img" :alt="item.alt">
    </div>
  </div>
</template>

<script setup>
import { useFetch } from '../composables/useFetch.js'

const { data: hero, loading } = useFetch('/api/hero', { quotes: [] })
</script>
