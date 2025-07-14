<template>
  <div v-if="isVisible" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white p-6 rounded-lg shadow-xl max-w-sm w-full">
      <h3 class="text-xl font-serif text-light-text mb-4">{{ title }}</h3>
      <p class="text-neutral-700 mb-6 font-sans">{{ message }}</p>
      <div class="flex justify-end space-x-4">
        <button @click="cancel" class="align-middle select-none font-sans font-bold text-center uppercase transition-all text-xs py-2 px-4 rounded-lg bg-neutral-200 text-neutral-700 hover:bg-neutral-300 focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none">
          Cancel
        </button>
        <button @click="confirm" class="align-middle select-none font-sans font-bold text-center uppercase transition-all text-xs py-2 px-4 rounded-lg bg-red-500 text-dark-text shadow-md hover:shadow-lg focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none">
          Confirm
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ConfirmationModal',
  props: {
    isVisible: {
      type: Boolean,
      default: false,
    },
    title: {
      type: String,
      default: 'Confirm Action',
    },
    message: {
      type: String,
      default: 'Are you sure you want to proceed?',
    },
  },
  methods: {
    confirm() {
      this.$emit('confirm');
    },
    cancel() {
      this.$emit('cancel');
    },
  },
};
</script>
