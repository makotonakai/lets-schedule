<template>
  <div class="wrapper">
    <nav class="navbar" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <a class="navbar-item">
          <router-link to="/home">
            <img src="../../assets/logo.png" width="140" height="64">
          </router-link>
        </a>
      </div>

      <div class="navbar-end">
        <div class="navbar-item">
          <div class="buttons">
            <router-link to="/signup" class="button is-primary">
              <strong>Sign up</strong>
            </router-link>
          </div>
        </div>
      </div>
    </nav>
  </div>
</template>