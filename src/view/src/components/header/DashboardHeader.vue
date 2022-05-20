<template>
  <div class="wrapper">
    <nav class="navbar" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <a class="navbar-item">
          <router-link to="/">
            <img src="../../assets/logo.png" width="140" height="64" />
          </router-link>
        </a>
      </div>

      <a
        role="button"
        class="navbar-burger"
        aria-label="menu"
        aria-expanded="false"
        data-target="targetMenu"
      >
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
      </a>

      <div id="targetMenu" class="navbar-menu">
        <div class="navbar-start">
          <a class="navbar-item"> ミーティング新規作成 </a>

          <div class="navbar-item has-dropdown is-hoverable">
            <a class="navbar-link"> 主催ミーティング確認 </a>

            <div class="navbar-dropdown">
              <a class="navbar-item"> 確定済み </a>
              <hr class="navbar-divider" />
              <a class="navbar-item"> 未確定 </a>
            </div>
          </div>

          <div class="navbar-item has-dropdown is-hoverable">
            <a class="navbar-link"> 参加ミーティング確認 </a>

            <div class="navbar-dropdown">
              <a class="navbar-item"> 確定済み </a>
              <hr class="navbar-divider" />
              <a class="navbar-item"> 未対応 </a>
              <a class="navbar-item"> 返信済み </a>
            </div>
          </div>
        </div>
      </div>

      <div class="navbar-end">
        <div class="navbar-item">
          <div class="buttons">
            <router-link to="/" class="button is-primary">
              <strong>Log out</strong>
            </router-link>
          </div>
        </div>
      </div>
    </nav>
  </div>
</template>