package com.example.app_pos.presentation.viewmodel

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.app_pos.domain.entity.User
import com.example.app_pos.domain.entity.UserRole
import com.example.app_pos.domain.usecase.GetUsersUseCase
import com.example.app_pos.domain.usecase.GetUsersByRoleUseCase
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import javax.inject.Inject

@HiltViewModel
class UserViewModel @Inject constructor(
    private val getUsersUseCase: GetUsersUseCase,
    private val getUsersByRoleUseCase: GetUsersByRoleUseCase
) : ViewModel() {

    private val _uiState = MutableStateFlow(UserUiState())
    val uiState: StateFlow<UserUiState> = _uiState.asStateFlow()

    init {
        loadUsers()
    }

    fun loadUsers() {
        viewModelScope.launch {
            _uiState.value = _uiState.value.copy(isLoading = true, error = null)
            
            getUsersUseCase().fold(
                onSuccess = { users ->
                    _uiState.value = _uiState.value.copy(
                        users = users,
                        isLoading = false
                    )
                },
                onFailure = { error ->
                    _uiState.value = _uiState.value.copy(
                        error = error.message ?: "Unknown error occurred",
                        isLoading = false
                    )
                }
            )
        }
    }

    fun loadUsersByRole(role: UserRole) {
        viewModelScope.launch {
            _uiState.value = _uiState.value.copy(isLoading = true, error = null)
            
            getUsersByRoleUseCase(role.name.lowercase()).fold(
                onSuccess = { users ->
                    _uiState.value = _uiState.value.copy(
                        users = users,
                        selectedRole = role,
                        isLoading = false
                    )
                },
                onFailure = { error ->
                    _uiState.value = _uiState.value.copy(
                        error = error.message ?: "Unknown error occurred",
                        isLoading = false
                    )
                }
            )
        }
    }

    fun clearError() {
        _uiState.value = _uiState.value.copy(error = null)
    }

    fun selectRole(role: UserRole?) {
        if (role == null) {
            loadUsers()
        } else {
            loadUsersByRole(role)
        }
    }
}

data class UserUiState(
    val users: List<User> = emptyList(),
    val selectedRole: UserRole? = null,
    val isLoading: Boolean = false,
    val error: String? = null
)