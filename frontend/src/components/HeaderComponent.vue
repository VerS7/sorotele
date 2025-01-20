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

#payment {
  margin-right: 4.5rem;
}
</style>

<template>
  <v-container class="base-width pb-0">
    <v-container fluid class="d-flex pb-0">
      <v-col>
        <v-img :width="330" src="/logo.jpg"></v-img>
      </v-col>
      <v-col class="d-flex align-end justify-end">
        <v-card max-width="240" class="pa-2 rounded-lg" elevation="1">
          <div id="contacts" class="d-flex flex-column justify-center">
            <div class="d-flex justify-end align-center">
              <v-icon :size="30" id="phone" class="mr-1">mdi-phone</v-icon>
              <a id="phone__number">8 (495) 99-111-08</a>
            </div>
            <a href="mailto:support@sorotele.com" id="support__email" class="link"
              >support@sorotele.com</a
            >
          </div>
        </v-card>
      </v-col>
    </v-container>
    <v-container class="d-flex align-center flex-column full-width" min-width="1100">
      <nav>
        <NavButtonComponent class="rounded-s-xl" title="Главная" link-to="/"> </NavButtonComponent>
        <NavButtonComponent title="Услуги" link-to="/service"></NavButtonComponent>
        <NavButtonComponent title="Тарифы" link-to="/tarifs"></NavButtonComponent>
        <NavButtonComponent title="Абонентам" link-to="/customers"></NavButtonComponent>
        <NavButtonComponent title="Заявка на подключение" link-to="/order"></NavButtonComponent>
        <NavButtonComponent title="Контакты" link-to="/contacts"></NavButtonComponent>
        <NavButtonComponent
          class="rounded-e-xl"
          title="Личный кабинет"
          link-to="/login"
          link-refered="/lk"
          icon="mdi-account-circle"
        ></NavButtonComponent>
      </nav>
      <RouterLink
        id="payment"
        class="link__reverse color-green mt-1 full-width text-end mt-2"
        to="/payment"
        >Пополнение счета онлайн</RouterLink
      >
    </v-container>
  </v-container>
</template>

<script setup lang="ts">
import NavButtonComponent from '@/components/NavButtonComponent.vue'
</script>
