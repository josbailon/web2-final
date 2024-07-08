package aparcamiento.AparcamientoXI.jwt;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import lombok.RequiredArgsConstructor;

@RestController
@RequestMapping("api/v1")
@RequiredArgsConstructor
public class jwtController {
	
	@PostMapping(value ="jwt")
	public String welcome()
	{
		return "Welcome form secure endpoint";
	}

}
