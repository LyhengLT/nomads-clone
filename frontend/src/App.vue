<template>

  <!-- Video Container -->
  <Hero v-if="heroVisible" @openPopup="showPopup = true"/>
  <!-- End of Video Container -->

  <!-- Three big news quotes -->
  <NewsQuotes v-if="!filterOpen"/>
  <!-- End of News Quotes-->

  <!-- Filter Sidebar -->
  <div class="lh-page-filter-sidebar">
    <FilterSidebar :isOpen="filterOpen"/>
    <div class="lh-page-main-content" :class="{ 'lh-filter-open': filterOpen }">
      <NewsQuotes v-if="filterOpen"/>
      <ViewContainer :filterOpen="filterOpen"
                     @toggleFilter="toggleFilter"
                     @openPopup="showPopup = true"/>
      <GridCards/>
    </div>
  </div>
  <!-- End of Filter side bar -->

  <!-- Footer and Popups -->
  <Footer @openPopup="showPopup = true"/>
  <AdsPopup :isOpen="showPopup" @close="showPopup = false"/>
  <!-- End of Footer and Popups -->
</template>

<script setup>
import {ref, onMounted} from 'vue'
import Hero from './components/Hero.vue'
import NewsQuotes from './components/NewsQuotes.vue'
import ViewContainer from './components/ViewContainer.vue'
import FilterSidebar from './components/FilterSidebar.vue'
import GridCards from './components/GridCards.vue'
import Footer from './components/Footer.vue'
import AdsPopup from './components/AdsPopup.vue'

const showPopup = ref(false)
const filterOpen = ref(false)
const heroVisible = ref(true)

function toggleFilter() {
  filterOpen.value = !filterOpen.value
  if (filterOpen.value) {
    heroVisible.value = false
    window.scrollTo(0, 0)
  }
}

onMounted(() => {
  if ('scrollRestoration' in history) {
    history.scrollRestoration = 'manual'
  }
  window.scrollTo(0, 0)

  setTimeout(() => {
    showPopup.value = true
  }, 1500)
})
</script>