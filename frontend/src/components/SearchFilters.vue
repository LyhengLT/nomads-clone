<template>
  <div class="search-filter-results" id="search-filters">
    <div v-for="(col, i) in filters.columns" :key="i" class="search-filter-results-column">
      <div v-for="item in col" :key="item.label"
           :data-label="item.label"
           :data-tooltip="item.tooltip"
           class="search-filter-result tooltip-set">
        {{ item.text }}
      </div>
    </div>
  </div>
</template>

<script setup>
import { useFetch } from '../composables/useFetch.js'

const { data: filters } = useFetch('/api/filters', { columns: [] })
</script>