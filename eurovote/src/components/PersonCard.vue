<template>
    <div class="person-card">
      <h3>{{ personData.Label }} ({{ personData.Identifier }})</h3>
      <p><strong>Given Name:</strong> {{ personData.GivenName }}</p>
      <p><strong>Family Name:</strong> {{ personData.FamilyName }}</p>
      <a :href="personData.About" target="_blank">Profile Link</a>
    </div>
  </template>
  
  <script>
  export default {
    name: 'PersonCard',
    props: {
      personData: Object
    }
  }
  </script>
  
  <style>
  .person-card {
    border: 1px solid #ddd;
    border-radius: 8px;
    padding: 16px;
    margin: 16px;
    width: 300px;
  }
  .person-card h3 {
    margin-top: 0;
  }
  </style>
  