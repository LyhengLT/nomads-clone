<template>
  <div class="lh-filter-sidebar" v-if="isOpen" @click="toggleOption">

    <!-- WHAT -->
    <div class="lh-header-filter">WHAT</div>
    <div class="lh-row-filter lh-three">
      <span class="lh-tooltip lh-two-btn">🍦 Cold Now</span>
      <span class="lh-tooltip lh-two-btn">🌤️ Mild now</span>
      <span class="lh-tooltip lh-two-btn">☀️ Warm now</span>
    </div>
    <div class="lh-row-filter lh-three">
      <span class="lh-tooltip lh-two-btn">💵 &lt;US$1k/mo</span>
      <span class="lh-tooltip lh-two-btn">💸 &lt;US$2k/mo</span>
      <span class="lh-tooltip lh-two-btn">💰 &lt;US$3k/mo</span>
    </div>
    <div class="lh-row-filter lh-two">
      <span class="lh-tooltip lh-two-btn">👮‍ Safe</span>
      <span class="lh-tooltip lh-two-btn">📡 Fast internet</span>
    </div>
    <div class="lh-row-filter lh-two">
      <span class="lh-tooltip lh-two-btn">💨 Clean air now</span>
      <span class="lh-tooltip lh-two-btn">👍 Liked by members</span>
    </div>
    <div class="lh-row-filter lh-two">
      <span class="lh-tooltip lh-two-btn">🔥 Popular now</span>
      <span class="lh-tooltip lh-two-btn">📈 Growing in nomads</span>
    </div>
    <div class="lh-row-filter lh-two lh-stick">
      <span class="lh-tooltip lh-two-btn">🏅 Top ranked</span>
      <span class="lh-tooltip lh-two-btn">💎 Hidden gem</span>
    </div>
    <div class="lh-row-filter lh-two">
      <span class="lh-tooltip lh-two-btn">✨ You haven't been</span>
      <span class="lh-tooltip lh-two-btn">✨ For you</span>
    </div>
    <div class="lh-row-filter lh-two">
      <span class="lh-tooltip lh-two-btn-one-box">🇰🇭 Near Cambodia</span>
    </div>

    <!-- WHERE -->
    <div class="lh-header-filter">WHERE</div>
    <div class="lh-where-grid">
      <h3 class="lh-tooltip">🌎 North America</h3>
      <h3 class="lh-tooltip">💃 Latin America</h3>
      <h3 class="lh-tooltip">🇪🇺 Europe</h3>
      <h3 class="lh-tooltip">🌍 Africa</h3>
      <h3 class="lh-tooltip">🕌 Middle East</h3>
      <h3 class="lh-tooltip">⛩️ Asia</h3>
      <h3 class="lh-tooltip">🏄 Oceania</h3>
      <h3 class="lh-tooltip">🛰️ Space</h3>
    </div>
    <div class="lh-row-filter lh-two">
      <span class="lh-tooltip lh-two-btn">🇺🇸 United States</span>
      <span class="lh-tooltip lh-two-btn">🏝 Caribbean</span>
    </div>
    <div class="lh-row-filter lh-two">
      <span class="lh-tooltip lh-two-btn">🇪🇺 European Union</span>
      <span class="lh-tooltip lh-two-btn">🇪🇺 Not in Schengen</span>
    </div>
    <div class="lh-row-filter lh-two">
      <span class="lh-tooltip lh-two-btn">🤝 Easy to make friends</span>
      <span class="lh-tooltip lh-two-btn">❤️ Great for dating</span>
    </div>

    <!-- WHEN -->
    <div class="lh-header-filter">WHEN</div>
    <div class="lh-row-filter lh-two lh-stick">
      <span class="lh-tooltip lh-two-btn">☃️ In the winter</span>
      <span class="lh-tooltip lh-two-btn">♻️ All year round</span>
    </div>
    <div class="lh-where-grid lh-when">
      <h3>Jan</h3><h3>Feb</h3><h3>Mar</h3><h3>Apr</h3>
      <h3>May</h3><h3>Jun</h3><h3>Jul</h3><h3>Aug</h3>
      <h3>Sep</h3><h3>Oct</h3><h3>Nov</h3><h3>Dec</h3>
    </div>

    <!-- SCORE -->
    <div class="lh-header-filter">SCORE</div>
    <div class="lh-row-filter lh-two">
      <span class="lh-tooltip lh-two-btn">🤩 Exceptional (4.75+)</span>
      <span class="lh-tooltip lh-two-btn">😍 Very good (4.5+)</span>
    </div>
    <div class="lh-row-filter lh-two">
      <span class="lh-tooltip lh-two-btn">👍 Good (4.25+)</span>
      <span class="lh-tooltip lh-two-btn">😐 Okay (3+)</span>
    </div>

    <!-- VISA FREE FOR -->
    <div class="lh-header-filter">VISA FREE FOR</div>
    <select class="lh-row-filter lh-passport">
      <option value="">Select your passport</option>
      <option v-for="opt in sidebar.passportOptions" :key="opt">{{ opt }}</option>
    </select>
    <div class="lh-plus-filter">+</div>
    <select class="lh-row-filter lh-passport">
      <option value="">Select your partner's passport</option>
      <option v-for="opt in sidebar.passportOptions" :key="opt">{{ opt }}</option>
    </select>
    <div class="lh-plus-filter">+</div>
    <select class="lh-row-filter lh-passport">
      <option value="">Select your friend's passport</option>
      <option v-for="opt in sidebar.passportOptions" :key="opt">{{ opt }}</option>
    </select>

    <!-- HOW -->
    <div class="lh-header-filter">HOW</div>
    <div class="lh-row-filter lh-two">
      <span class="lh-tooltip lh-two-btn">👮‍ Not Humid</span>
      <span class="lh-tooltip lh-two-btn">📡 Not rainy now</span>
    </div>
    <div class="lh-row-filter lh-two">
      <span class="lh-tooltip lh-two-btn">💨 Clean air now</span>
      <span class="lh-tooltip lh-two-btn">👍 Liked by members</span>
    </div>
    <div class="lh-row-filter lh-two">
      <span class="lh-tooltip lh-two-btn">🔥 Popular now</span>
      <span class="lh-tooltip lh-two-btn">📈 Growing in nomads</span>
    </div>
    <div class="lh-row-filter lh-two lh-stick">
      <span class="lh-tooltip lh-two-btn">🏅 Top ranked</span>
      <span class="lh-tooltip lh-two-btn">💎 Hidden gem</span>
    </div>
    <div class="lh-row-filter lh-two">
      <span class="lh-tooltip lh-two-btn">✨ You haven't been</span>
      <span class="lh-tooltip lh-two-btn">✨ For you</span>
    </div>

  </div>
</template>

<script setup>
import { useFetch } from '../composables/useFetch.js'

defineProps(['isOpen'])

const { data: sidebar } = useFetch('/api/sidebar', { passportOptions: [] })

function toggleOption(event) {
  const option = event.target.closest('.lh-two-btn, .lh-two-btn-one-box, .lh-where-grid h3')
  if (!option) return
  option.classList.toggle('lh-active')
}
</script>
