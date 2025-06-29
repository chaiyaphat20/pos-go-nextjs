package com.example.app_pos.presentation.viewmodel;

import com.example.app_pos.domain.usecase.GetUsersByRoleUseCase;
import com.example.app_pos.domain.usecase.GetUsersUseCase;
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
public final class UserViewModel_Factory implements Factory<UserViewModel> {
  private final Provider<GetUsersUseCase> getUsersUseCaseProvider;

  private final Provider<GetUsersByRoleUseCase> getUsersByRoleUseCaseProvider;

  public UserViewModel_Factory(Provider<GetUsersUseCase> getUsersUseCaseProvider,
      Provider<GetUsersByRoleUseCase> getUsersByRoleUseCaseProvider) {
    this.getUsersUseCaseProvider = getUsersUseCaseProvider;
    this.getUsersByRoleUseCaseProvider = getUsersByRoleUseCaseProvider;
  }

  @Override
  public UserViewModel get() {
    return newInstance(getUsersUseCaseProvider.get(), getUsersByRoleUseCaseProvider.get());
  }

  public static UserViewModel_Factory create(Provider<GetUsersUseCase> getUsersUseCaseProvider,
      Provider<GetUsersByRoleUseCase> getUsersByRoleUseCaseProvider) {
    return new UserViewModel_Factory(getUsersUseCaseProvider, getUsersByRoleUseCaseProvider);
  }

  public static UserViewModel newInstance(GetUsersUseCase getUsersUseCase,
      GetUsersByRoleUseCase getUsersByRoleUseCase) {
    return new UserViewModel(getUsersUseCase, getUsersByRoleUseCase);
  }
}
