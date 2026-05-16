<template>
  <div class="lh-ads-pop" v-if="isOpen" @click="closeOutside">
    <div class="lh-ads-pop-container" @click.stop>
<!--<button class="lh-ads-pop-close" @click="$emit('close')">✕</button>-->

      <!-- Left -->
      <div class="lh-ads-left">
        <div class="lh-stars">
          <span>⭐⭐⭐⭐⭐</span>
          <span class="lh-rating">Customer rating 9.0 | 8,350 reviews</span>
        </div>
        <template v-for="f in meetups.popupFeatures" :key="f.text">
          <div class="lh-popup-info lh-text-info">
            <span>✅ <span v-html="f.text"></span></span>
          </div>
          <div v-if="f.img" class="lh-popup-img">
            <img :src="f.img" :alt="f.imgAlt">
          </div>
        </template>
        <div class="lh-paragraph">
          <span>By signing up, you agree to our</span>
          <span class="lh-paragraph-color">terms of service and community guidelines.</span>
        </div>
      </div>

      <!-- Right -->
      <div class="lh-ads-right">
        <div class="lh-video">
          <img :src="meetups.popupRight.thumbnail" alt="Explore the world">
          <svg fill="#fff" viewBox="0 0 60 60" xmlns="http://www.w3.org/2000/svg">
            <path
                d="M30,0C13.458,0,0,13.458,0,30s13.458,30,30,30s30-13.458,30-30S46.542,0,30,0z M45.563,30.826l-22,15C23.394,45.941,23.197,46,23,46c-0.16,0-0.321-0.038-0.467-0.116C22.205,45.711,22,45.371,22,45V15c0-0.371,0.205-0.711,0.533-0.884c0.328-0.174,0.724-0.15,1.031,0.058l22,15C45.836,29.36,46,29.669,46,30S45.836,30.64,45.563,30.826z"/>
          </svg>
        </div>
        <div class="lh-members-joined">{{ meetups.popupRight.membersJoined }}</div>
        <div class="lh-email-box">
          <input type="email" :placeholder="meetups.popupRight.emailPlaceholder">
        </div>
        <div class="lh-gender">
          <select>
            <option v-for="opt in meetups.popupRight.genderOptions" :key="opt.label"
                    :selected="opt.selected">{{ opt.label }}
            </option>
          </select>
        </div>
        <div class="lh-join-nomads">
          <a :href="meetups.popupRight.joinHref">{{ meetups.popupRight.joinBtn }}</a>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup>
import { useFetch } from '../composables/useFetch.js'

const props = defineProps(['isOpen'])
const emit = defineEmits(['close'])

const { data: meetups } = useFetch('/api/meetups', {
  popupFeatures: [],
  popupRight: { thumbnail: '', membersJoined: '', emailPlaceholder: '', genderOptions: [], joinBtn: '', joinHref: '#' }
})

function closeOutside() {
  emit('close')
}
</script>