package com.parking.parqueadero.auth;

import org.springframework.security.core.userdetails.User;
import org.springframework.stereotype.Service;

import com.parking.parqueadero.user.UserRepository;
import com.parking.parqueadero.user.*;

import lombok.RequiredArgsConstructor;

@Service
@RequiredArgsConstructor
public class AuthenticationService {
	
	private final UserRepository repository = null;

	public AuthenticationResponse register(RegisterRequest request) {
		var user = User.builder()
				.build()
		return null;
	}

	public AuthenticationResponse authenticate(AuthenticationRequest request) {
		// TODO Auto-generated method stub
		return null;
	}

}
