<template>
  <div class="grid-cards" id="gridCards" v-show="!loading">
    <template v-for="item in cities" :key="item.id || item.type">

      <!-- City Card -->
      <div v-if="!item.type" class="grid-card">
        <div v-if="item.popular" class="grid-card-popular">Popular</div>
        <img :src="item.img" :alt="item.name" class="grid-card-img">
        <div class="grid-card-hover">
          <svg class="hover-heart" aria-hidden="true" width="25" focusable="false" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
            <path fill="none" stroke="white" stroke-width="30" d="M462.3 62.6C407.5 15.9 326 24.3 275.7 76.2L256 96.5l-19.7-20.3C186.1 24.3 104.5 15.9 49.7 62.6c-62.8 53.6-66.1 149.8-9.9 207.9l193.5 199.8c12.5 12.9 32.8 12.9 45.3 0l193.5-199.8c56.3-58.1 53-154.3-9.8-207.9z"></path>
          </svg>
          <span class="hover-close" data-tooltip="Not interested">✕</span>
          <div v-for="bar in barLabels" :key="bar.key" class="hover-row">
            <span class="hover-row-label">{{ bar.emoji }} {{ bar.label }}</span>
            <div class="hover-bar-track">
              <div :class="['hover-bar', item.bars[bar.key]]"></div>
            </div>
          </div>
        </div>
        <div class="grid-card-top">
          <span class="grid-card-rank">{{ item.rank }}</span>
          <span class="grid-card-wifi">
            <img :src="base + '/images/cities/wifi.svg'" alt="wifi" width="18" height="18">
            <span class="wifi-info">
              <span class="wifi-number">{{ item.wifi }}</span>
              <span class="wifi-mbps">Mbps</span>
            </span>
          </span>
        </div>
        <div class="grid-card-content">
          <div class="grid-card-name"><h3>{{ item.name }}</h3><p>{{ item.country }}</p></div>
          <div class="grid-card-bottom">
            <div class="grid-card-weather">
              {{ item.weather.icon }}
              <span class="weather-temp-group">
                <span class="feels">FEELS {{ item.weather.feels }}°</span>
                <span class="temp">{{ item.weather.temp }}°</span>
              </span>
              <template v-if="item.weather.humid"> {{ item.weather.humid }}</template>
              <template v-if="item.weather.aqi > 0">
                <span class="weather-aqi-group">
                  <span class="aqi-1">AQI</span>
                  <span class="aqi-value-1">{{ item.weather.aqi }}</span>
                </span>
              </template>
              <template v-if="item.weather.aqiIcon"> {{ item.weather.aqiIcon }}</template>
            </div>
            <div class="grid-card-cost">
              <span class="cost">{{ item.cost }}</span>
              <span class="cost-label">FOR A NOMAD</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Ad One -->
      <div v-else-if="item.type === 'ad-one'" class="sidebar-boxs ad-one">
        <div class="sidebar-img"><img :src="base + '/images/cities/safetywing.jpg'" alt="ad"></div>
        <div class="sidebar-hero-text">
          <h3>Global travel and health for nomads, starting at $2/day</h3>
          <a href="#">Get insured</a>
          <p class="ad">Ad</p>
        </div>
      </div>

      <!-- Ad Two -->
      <div v-else-if="item.type === 'ad-two'" class="sidebar-boxs ad-two">
        <div class="sidebar-img"><img :src="base + '/images/cities/residency.jpg'" alt="residency ad"></div>
        <div class="sidebar-hero-text">
          <h3>Get legal residency in 🇪🇸 Spain or 🇵🇹 Portugal. Passport and more.</h3>
          <a href="#">Get residency</a>
          <p class="ad">Ad</p>
        </div>
      </div>

      <!-- Meetups -->
      <div v-else-if="item.type === 'meetups'" class="sidebar-boxs meetups">
        <div class="sidebar-meetup-header meet-up-bottom-border">🥥 Next meetups (18/mo)</div>
        <div class="schedule-meetup">
          <div v-for="m in meetups.meetups" :key="m.label"
               class="meet-up-day meet-up-bottom-border">
            <div class="profile-meetup">
              <img :src="base + '/images/profiles/profile-' + m.avatar + '.png'" alt="avatar">
              <div class="sechudal-with-lable">{{ m.label }}
                <span class="meet-up-lable">{{ m.rsvp }} RSVPS</span>
              </div>
            </div>
            <div v-if="m.attendees.length" class="two-profile">
              <img v-for="n in m.attendees" :key="n"
                   :src="base + '/images/profiles/profile-' + n + '.png'" :alt="'p' + n">
            </div>
          </div>
        </div>
        <div class="see-meeting">See upcoming meetups →</div>
      </div>

      <!-- Traveling -->
      <div v-else-if="item.type === 'traveling'" class="sidebar-boxs traveling">
        <div class="sidebar-meetup-header meet-up-bottom-border travlering">🛩️ Traveling now (17)</div>
        <div class="img-travler">
          <div v-for="n in meetups.traveling" :key="'t-' + n" class="travler-images">
            <img :src="base + '/images/profiles/profile-' + n + '.png'" :alt="'p' + n">
            <div class="travler-popup-img">
              <p>Hidden for non-members due to privacy</p>
              <img :src="base + '/images/profiles/profile-' + n + '.png'" alt="blurred profile">
            </div>
          </div>
        </div>
      </div>

      <!-- New Members -->
      <div v-else-if="item.type === 'new-member'" class="sidebar-boxs new-member">
        <div class="sidebar-meetup-header meet-up-bottom-border travlering">👋 New members (653/mo)</div>
        <div class="img-travler">
          <div v-for="n in meetups.newMembers" :key="'nm-' + n" class="travler-images">
            <img :src="base + '/images/profiles/profile-' + n + '.png'" :alt="'p' + n">
            <div class="travler-popup-img">
              <p>Hidden for non-members due to privacy</p>
              <img :src="base + '/images/profiles/profile-' + n + '.png'" alt="blurred profile">
            </div>
          </div>
        </div>
      </div>

      <!-- Today Pick -->
      <div v-else-if="item.type === 'today-pick'" class="sidebar-boxs today-pick">
        <div class="grid-card-popular">Today's pick</div>
        <h3>Visit Places Visa Free for Citizens of Brazil</h3>
        <p>210x viewed</p>
      </div>

    </template>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useFetch } from '../composables/useFetch.js'

const base = import.meta.env.BASE_URL.replace(/\/$/, '')

const { data: cities, loading: citiesLoading } = useFetch('/api/cities', [])
const { data: meetups, loading: meetupsLoading } = useFetch('/api/meetups', {
  meetups: [], traveling: [], newMembers: []
})

const loading = computed(() => citiesLoading.value || meetupsLoading.value)

const barLabels = [
  { key: 'overall',  emoji: '⭐', label: 'Overall' },
  { key: 'cost',     emoji: '💵', label: 'Cost' },
  { key: 'internet', emoji: '📡', label: 'Internet' },
  { key: 'liked',    emoji: '👍', label: 'Liked' },
  { key: 'safety',   emoji: '👮', label: 'Safety' },
]
</script>