<script setup lang="ts">
import { ref, type PropType, watch } from 'vue'
import type { ComparisonType } from '../functions'

const emit = defineEmits(['update:modelValue'])

const props = defineProps({
    modelValue: {
        type: Object as PropType<ComparisonType>,
        required: false
    }
})

const firstInput = ref(props.modelValue?.leftValue || '')
const selectedOption = ref(props.modelValue?.operator || '')
const lastInput = ref(props.modelValue?.rightValue || '')

const options = ['Equals', 'Not Equals'] // Replace with your actual options

watch(
    [firstInput, selectedOption, lastInput],
    ([left, operator, right]) => {
        emit('update:modelValue', {
            leftValue: left,
            rightValue: right,
            operator: operator
        })
    },
    { immediate: true }
)

</script>

<template>
  <div class="comparison-container">
    <div class="input-column">
      <input 
        v-model="firstInput"
        type="text"
        placeholder="Enter first value"
        class="text-input"
      />
    </div>

    <div class="input-column">
      <select 
        v-model="selectedOption"
        class="dropdown"
      >
        <option value="" disabled>Select an option</option>
        <option 
          v-for="option in options" 
          :key="option" 
          :value="option"
        >
          {{ option }}
        </option>
      </select>
    </div>

    <div class="input-column">
      <input 
        v-model="lastInput"
        type="text"
        placeholder="Enter last value"
        class="text-input"
      />
    </div>
  </div>
</template>

<style scoped>
.comparison-container {
  display: grid;
  grid-template-columns: minmax(150px, 1fr) minmax(150px, 1fr) minmax(150px, 1fr);
  gap: 20px;
  padding: 20px;
  min-width: 500px;
}

.input-column {
  width: 100%;
}

.text-input, .dropdown {
  width: 100%;
  padding: 12px 16px;
  background: #2d2d2d;
  border: 1px solid #404040;
  border-radius: 6px;
  color: #fff;
  font-size: 1rem;
  transition: all 0.2s ease;
}

.text-input::placeholder {
  color: #666;
}

.dropdown {
  cursor: pointer;
  appearance: none;
  background-repeat: no-repeat;
  background-position: right 12px center;
  padding-right: 36px;
}

.dropdown option {
  background: #2d2d2d;
  color: #fff;
}
</style>