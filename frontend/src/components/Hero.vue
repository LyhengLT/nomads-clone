<template>
  <section class="hero" v-show="!loading">
    <video class="hero-video" autoplay muted loop playsinline>
      <source :src="base + '/videos/nomad-vdo-hd.mp4'" type="video/mp4">
    </video>
    <div class="hero-overlay"></div>

    <div class="hero-inner">
      <!-- Left -->
      <div class="hero-left">
        <div class="laurel">
          <span class="laurel-text">{{ hero.laurel.text }}</span>
          <span class="laurel-stars">
            <img v-for="i in hero.laurel.stars" :key="i" :src="base + '/images/star.svg'" alt="star">
          </span>
          <span class="laurel-since">{{ hero.laurel.since }}</span>
          <img class="laurel-img" :src="hero.laurel.img" alt="laurel">
        </div>
        <h3 class="hero-title">{{ hero.heroTitle }}</h3>
        <p class="hero-subtitle-text">{{ hero.heroSubtitle }}</p>
        <div class="hero-profiles">
          <img v-for="(src, i) in hero.profiles" :key="i"
               :src="src" :alt="'profile' + (i+1)" class="hero-profile">
        </div>
        <div v-for="b in hero.benefits" :key="b.text" class="benefit">
          {{ b.emoji }} <a :href="b.href">{{ b.text }}</a>
        </div>
      </div>

      <!-- Right -->
      <div class="hero-right">
        <div class="cta-box">
          <div class="intro-video-box">
            <img :src="hero.ctaBox.thumbnail" alt="Explore the world">
            <svg class="play-icon" fill="#fff" viewBox="0 0 60 60" xmlns="http://www.w3.org/2000/svg">
              <path d="M30,0C13.458,0,0,13.458,0,30s13.458,30,30,30s30-13.458,30-30S46.542,0,30,0z M45.563,30.826l-22,15C23.394,45.941,23.197,46,23,46c-0.16,0-0.321-0.038-0.467-0.116C22.205,45.711,22,45.371,22,45V15c0-0.371,0.205-0.711,0.533-0.884c0.328-0.174,0.724-0.15,1.031,0.058l22,15C45.836,29.36,46,29.669,46,30S45.836,30.64,45.563,30.826z"/>
            </svg>
          </div>
          <p><input type="email" class="email" :placeholder="hero.ctaBox.emailPlaceholder"></p>
          <span class="btn-join join-nomads" @click="$emit('openPopup')">{{ hero.ctaBox.joinBtn }}</span>
          <p class="already-have-account">{{ hero.ctaBox.alreadyAccount }}</p>
        </div>
      </div>
    </div>

    <!-- Wave -->
    <div class="hero-wave">
      <svg viewBox="0 0 1440 120" preserveAspectRatio="none" xmlns="http://www.w3.org/2000/svg">
        <path d="M1440,21.2101911 L1440,120 L0,120 L0,21.2101911 C120,35.0700637 240,42 360,42 C480,42 600,35.0700637 720,21.2101911 C808.32779,12.416393 874.573633,6.87702029 918.737528,4.59207306 C972.491685,1.8109458 1026.24584,0.420382166 1080,0.420382166 C1200,0.420382166 1320,7.35031847 1440,21.2101911 Z"></path>
      </svg>
    </div>

    <!-- News Logos -->
    <div class="news-logo">
      <span class="news-seen">as seen on</span>
      <img v-for="logo in hero.newsLogos" :key="logo.alt"
           :src="logo.src" :alt="logo.alt">
    </div>
  </section>
</template>

<script setup>
import { useFetch } from '../composables/useFetch.js'

defineEmits(['openPopup'])

const base = import.meta.env.BASE_URL.replace(/\/$/, '')

const { data: hero, loading } = useFetch('/api/hero', {
  laurel: { text: '', stars: 0, since: '', img: '' },
  heroTitle: '',
  heroSubtitle: '',
  profiles: [],
  benefits: [],
  newsLogos: [],
  ctaBox: { thumbnail: '', emailPlaceholder: '', joinBtn: '', alreadyAccount: '' },
  quotes: []
})
</script>