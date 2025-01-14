<style scoped>
#contacts {
  color: #676153;
  text-align: end;
}

#phone__number {
  font-size: 1.5rem;
}

#support__email {
  font-size: 1.17rem;
}

#phone {
  color: #02d31e;
  justify-content: center;
  align-items: center;
}
</style>

<template>
  <v-container class="base-width">
    <v-container fluid class="d-flex">
      <v-col>
        <v-img :width="330" src="src/assets/logo.jpg"></v-img>
      </v-col>
      <v-col>
        <div id="contacts" class="d-flex flex-column justify-center">
          <div class="d-flex justify-end align-center">
            <v-icon :size="30" id="phone" class="mr-1">mdi-phone</v-icon>
            <a id="phone__number">8 (495) 99-111-08</a>
          </div>
          <a href="" id="support__email" class="link">support@sorotele.com</a>
        </div>
      </v-col>
    </v-container>
    <v-container class="d-flex align-center flex-column full-width">
      <nav>
        <NavButtonComponent title="Главная" link-to="/"> </NavButtonComponent>
        <NavButtonComponent title="Услуги" link-to="/service"></NavButtonComponent>
        <NavButtonComponent title="Тарифы" link-to="/tarifs"></NavButtonComponent>
        <NavButtonComponent title="Абонентам" link-to="/customers"></NavButtonComponent>
        <NavButtonComponent title="Заявка на подключение" link-to="/order"></NavButtonComponent>
        <NavButtonComponent title="Контакты" link-to="/contacts"></NavButtonComponent>
        <NavButtonComponent
          title="Личный кабинет"
          link-to="/lk"
          icon="mdi-account-circle"
        ></NavButtonComponent>
      </nav>
      <a class="link__reverse color-green mt-1 mr-1 full-width text-end mr-15 mt-2" href="#"
        >Пополнение счета онлайн</a
      >
    </v-container>
  </v-container>
</template>

<script setup lang="ts">
import NavButtonComponent from '@/components/NavButtonComponent.vue'
</script>
