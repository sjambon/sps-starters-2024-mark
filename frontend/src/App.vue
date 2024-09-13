<template>
  <div id="app">
    <h1>MARK: Campaign Manager</h1>
    <p>Welcome to AE!</p>
    <p>I am MARK, your campaign manager</p>
    <div v-if="loading" class="loading">Loading campaigns...</div>
    <div v-else-if="error" class="error">Error: {{ error }}</div>
    <div v-else class="campaigns-section">
      <h2>Campaigns</h2>
      <div class="campaign-list">
        <div class="campaign-card" v-for="campaign in campaigns" :key="campaign.id">
          <img :src="'https://ui-avatars.com/api/?name=' + encodeURIComponent(campaign.name)" alt="Campaign Icon"
            class="campaign-icon" />
          <div class="campaign-details">
            <h3>{{ campaign.name }}</h3>
            <p>{{ campaign.description }}</p>
          </div>
        </div>
      </div>
    </div>
    <div class="create-campaign">
      <h2>Create New Campaign</h2>
      <form @submit.prevent="createCampaign" class="campaign-form">
        <input v-model="newCampaign.name" placeholder="Campaign Name" required class="input-field" />
        <input v-model="newCampaign.description" placeholder="Campaign Description" required class="input-field" />
        <button type="submit" class="create-btn">Create Campaign</button>
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
    const loading = ref(true)
    const error = ref(null)

    const fetchCampaigns = async () => {
      try {
        loading.value = true
        const response = await axios.get(`${process.env.VUE_APP_API_URL}/campaigns`)
        campaigns.value = response.data
        loading.value = false
      } catch (err) {
        error.value = err.message
        loading.value = false
        console.error('Error fetching campaigns:', err)
      }
    }

    const createCampaign = async () => {
      try {
        await axios.post(`${process.env.VUE_APP_API_URL}/campaigns`, newCampaign.value)
        newCampaign.value = { name: '', description: '' }
        await fetchCampaigns()
      } catch (err) {
        error.value = err.message
        console.error('Error creating campaign:', err)
      }
    }

    onMounted(fetchCampaigns)

    return {
      campaigns,
      newCampaign,
      createCampaign,
      loading,
      error
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.loading,
.error {
  font-size: 18px;
  color: #e74c3c;
  margin: 20px;
}

.campaigns-section {
  margin: 20px 0;
}

.campaign-list {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 20px;
}

.campaign-card {
  background-color: #ecf0f1;
  border-radius: 8px;
  padding: 20px;
  width: 250px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  text-align: left;
}

.campaign-icon {
  border-radius: 50%;
  width: 60px;
  height: 60px;
  display: block;
  margin-bottom: 10px;
}

.campaign-details h3 {
  margin: 0 0 5px;
  font-size: 18px;
}

.campaign-details p {
  margin: 0;
  font-size: 14px;
  color: #7f8c8d;
}

.create-campaign {
  margin: 40px 0;
}

.campaign-form {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
}

.input-field {
  padding: 10px;
  font-size: 16px;
  width: 300px;
  border: 1px solid #bdc3c7;
  border-radius: 4px;
}

.create-btn {
  background-color: #3498db;
  color: white;
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
}

.create-btn:hover {
  background-color: #2980b9;
}
</style>
