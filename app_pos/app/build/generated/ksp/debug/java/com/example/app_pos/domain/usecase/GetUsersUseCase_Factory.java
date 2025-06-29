package com.example.app_pos.domain.usecase;

import com.example.app_pos.domain.repository.UserRepository;
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
public final class GetUsersUseCase_Factory implements Factory<GetUsersUseCase> {
  private final Provider<UserRepository> repositoryProvider;

  public GetUsersUseCase_Factory(Provider<UserRepository> repositoryProvider) {
    this.repositoryProvider = repositoryProvider;
  }

  @Override
  public GetUsersUseCase get() {
    return newInstance(repositoryProvider.get());
  }

  public static GetUsersUseCase_Factory create(Provider<UserRepository> repositoryProvider) {
    return new GetUsersUseCase_Factory(repositoryProvider);
  }

  public static GetUsersUseCase newInstance(UserRepository repository) {
    return new GetUsersUseCase(repository);
  }
}
