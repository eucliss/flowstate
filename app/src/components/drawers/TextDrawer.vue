<script setup lang="ts">
import { ref } from 'vue'

// Define props and emits for v-model
const props = defineProps({
  content: {
    type: String,
    required: true
  }
})

const emit = defineEmits(['update:content'])

// Function to handle content changes
const updateContent = (newValue: string) => {
  emit('update:content', newValue)
}
</script>

<template>
  <div class="text-drawer">
    <textarea
      :value="content"
      @input="updateContent(($event.target as HTMLTextAreaElement).value)"
      placeholder="Enter your text here..."
      class="text-input"
      rows="15"
    ></textarea>
  </div>
</template>

<style scoped>
.text-drawer {
  margin-bottom: 24px;
  width: 100%;
  display: block;
  min-height: 300px;
  background-color: #1a1a1a;
  padding: 16px;
}

.text-input {
  width: 100%;
  padding: 12px 16px;
  background: #2d2d2d;
  border: 1px solid #404040;
  border-radius: 6px;
  color: #fff;
  font-size: 1rem;
  transition: all 0.2s ease;
  resize: vertical;
  min-height: 300px;
  font-family: inherit;
  line-height: 1.5;
  display: block;
}

.text-input:focus {
  outline: none;
  border-color: #0066cc;
  box-shadow: 0 0 0 2px rgba(0, 102, 204, 0.2);
}

.text-input::placeholder {
  color: #666;
}
</style>
