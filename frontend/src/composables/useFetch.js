import { ref, onMounted } from 'vue'

export function useFetch(url, defaultValue) {
  const data = ref(defaultValue)
  const loading = ref(true)

  onMounted(async () => {
    try {
      const res = await fetch(url)
      if (res.ok) data.value = await res.json()
      else console.error(`[useFetch] ${url} → ${res.status}`)
    } catch (err) {
      console.error(`[useFetch] ${url} failed:`, err)
    } finally {
      loading.value = false
    }
  })

  return { data, loading }
}
