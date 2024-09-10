<template>
  <div id="app">
    <h1>MARK: Campaign Manager</h1>
    <div class="campaigns">
      <h2>Campaigns</h2>
      <ul>
        <li v-for="campaign in campaigns" :key="campaign.id">
          {{ campaign.name }}: {{ campaign.description }}
        </li>
      </ul>
    </div>
    <div class="create-campaign">
      <h2>Create New Campaign</h2>
      <form @submit.prevent="createCampaign">
        <input v-model="newCampaign.name" placeholder="Campaign Name" required>
        <input v-model="newCampaign.description" placeholder="Campaign Description" required>
        <button type="submit">Create Campaign</button>
      </form>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import axios from 'axios'

export default {
  name: 'App',
  setup() {
    const campaigns = ref([])
    const newCampaign = ref({ name: '', description: '' })

    const fetchCampaigns = async () => {
      try {
        const response = await axios.get(`${process.env.VUE_APP_API_URL}/campaigns`)
        campaigns.value = response.data
      } catch (error) {
        console.error('Error fetching campaigns:', error)
      }
    }

    const createCampaign = async () => {
      try {
        await axios.post(`${process.env.VUE_APP_API_URL}/campaigns`, newCampaign.value)
        newCampaign.value = { name: '', description: '' }
        await fetchCampaigns()
      } catch (error) {
        console.error('Error creating campaign:', error)
      }
    }

    onMounted(fetchCampaigns)

    return {
      campaigns,
      newCampaign,
      createCampaign
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>