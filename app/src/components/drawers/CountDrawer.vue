<script setup lang="ts">
import { ref } from 'vue'
import Comparison from '../Comparison.vue'
import type { ComparisonType } from '../../functions'
import type { PropType } from 'vue'
import { handleQueryTest } from '../../functions'
// Define props and emits for v-model
const props = defineProps({
  query: {
    type: String,
    required: true
  },
})

const emit = defineEmits(['update:query'])

const jsonData = ref('')
const count = ref(0)
// Function to handle query changes
const updateQuery = (newValue: string) => {
  emit('update:query', newValue)
}

const handleTest = async () => {
    const data = await handleQueryTest(props.query)
    count.value = data.length
    jsonData.value = JSON.stringify(data, null, 2)
}
</script>

<template>
  <div class="query-drawer">
    <textarea
      :value="query"
      @input="updateQuery(($event.target as HTMLTextAreaElement).value)"
      placeholder="Enter SQL query"
      class="sql-input"
      rows="4"
    ></textarea>
  </div>
  <div class="input-group">
        <div class="label-row">
        <label for="jsonData">JSON Data</label>
        <button class="test-button" @click="handleTest">
            <i class="fas fa-vial"></i>
            Test
        </button>
        </div>
        <textarea
            id="jsonData"
            v-model="jsonData"
            placeholder="Enter JSON data"
            class="json-input"
            rows="6"
        ></textarea>
  </div>
  <div class="input-group">
    <div class="number-display">
      <span class="count">{{ count }}</span>
    </div>
  </div>
</template>

<style scoped>
.sql-input {
  width: 100%;
  padding: 12px 16px;
  background: #2d2d2d;
  border: 1px solid #404040;
  border-radius: 6px;
  color: #fff;
  font-size: 1rem;
  transition: all 0.2s ease;
  resize: vertical;
  min-height: 150px;
  font-family: monospace;
}

.sql-input:focus {
  outline: none;
  border-color: #0066cc;
  box-shadow: 0 0 0 2px rgba(0, 102, 204, 0.2);
}

.sql-input::placeholder {
  color: #666;
}

.input-row {
  display: flex;
  gap: 20px;
  margin-bottom: 24px;
}

.input-row .input-group {
  margin-bottom: 0;
  flex: 1;
}

.input-group {
  margin-bottom: 24px;
}

.columns-container {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
  flex: 1;
  min-height: 0;
}

.column {
  background: #2d2d2d;
  border-radius: 8px;
  padding: 16px;
  border: 1px solid #404040;
  display: flex;
  flex-direction: column;
  height: 100%;
}

.column h4 {
  color: #fff;
  margin: 0 0 16px 0;
  font-size: 1rem;
  font-weight: 500;
}

.column-content {
  flex: 1;
}

/* Input styling */
.input-group label {
  display: block;
  color: #999;
  margin-bottom: 8px;
  font-size: 0.9rem;
}

input {
  width: 100%;
  padding: 12px 16px;
  background: #2d2d2d;
  border: 1px solid #404040;
  border-radius: 6px;
  color: #fff;
  font-size: 1rem;
  transition: all 0.2s ease;
}

input:focus {
  outline: none;
  border-color: #0066cc;
  box-shadow: 0 0 0 2px rgba(0, 102, 204, 0.2);
}

input::placeholder {
  color: #666;
}
.json-input {
  width: 100%;
  padding: 12px 16px;
  background: #2d2d2d;
  border: 1px solid #404040;
  border-radius: 6px;
  color: #fff;
  font-size: 1rem;
  transition: all 0.2s ease;
  resize: vertical;
  min-height: 200px;
  font-family: monospace;
}

.json-input:focus {
  outline: none;
  border-color: #0066cc;
  box-shadow: 0 0 0 2px rgba(0, 102, 204, 0.2);
}

.json-input::placeholder {
  color: #666;
}


.test-button {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 6px 12px;
  background: #2d2d2d;
  border: 1px solid #404040;
  border-radius: 4px;
  color: #fff;
  font-size: 0.9rem;
  cursor: pointer;
  transition: all 0.2s ease;
}

.test-button:hover {
  background: #404040;
}

.test-button i {
  font-size: 0.9rem;
}
.label-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.number-display {
  display: flex;
  justify-content: center;
  align-items: center;
  background: #2d2d2d;
  border: 1px solid #404040;
  border-radius: 6px;
  padding: 20px;
  margin: 20px 0;
}

.count {
  font-size: 2rem;
  color: #fff;
  font-weight: bold;
}
</style> 