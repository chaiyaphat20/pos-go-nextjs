package com.example.app_pos.presentation.viewmodel;

import com.example.app_pos.domain.usecase.CheckAuthUseCase;
import com.example.app_pos.domain.usecase.LoginUseCase;
import dagger.internal.DaggerGenerated;
import dagger.internal.Factory;
import dagger.internal.QualifierMetadata;
import dagger.internal.ScopeMetadata;
import javax.annotation.processing.Generated;
import javax.inject.Provider;

@ScopeMetadata
@QualifierMetadata
@DaggerGenerated
@Generated(
    value = "dagger.internal.codegen.ComponentProcessor",
    comments = "https://dagger.dev"
)
@SuppressWarnings({
    "unchecked",
    "rawtypes",
    "KotlinInternal",
    "KotlinInternalInJava"
})
public final class LoginViewModel_Factory implements Factory<LoginViewModel> {
  private final Provider<LoginUseCase> loginUseCaseProvider;

  private final Provider<CheckAuthUseCase> checkAuthUseCaseProvider;

  public LoginViewModel_Factory(Provider<LoginUseCase> loginUseCaseProvider,
      Provider<CheckAuthUseCase> checkAuthUseCaseProvider) {
    this.loginUseCaseProvider = loginUseCaseProvider;
    this.checkAuthUseCaseProvider = checkAuthUseCaseProvider;
  }

  @Override
  public LoginViewModel get() {
    return newInstance(loginUseCaseProvider.get(), checkAuthUseCaseProvider.get());
  }

  public static LoginViewModel_Factory create(Provider<LoginUseCase> loginUseCaseProvider,
      Provider<CheckAuthUseCase> checkAuthUseCaseProvider) {
    return new LoginViewModel_Factory(loginUseCaseProvider, checkAuthUseCaseProvider);
  }

  public static LoginViewModel newInstance(LoginUseCase loginUseCase,
      CheckAuthUseCase checkAuthUseCase) {
    return new LoginViewModel(loginUseCase, checkAuthUseCase);
  }
}
