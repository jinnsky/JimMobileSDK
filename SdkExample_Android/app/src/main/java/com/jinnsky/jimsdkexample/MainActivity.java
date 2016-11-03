package com.jinnsky.jimsdkexample;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;
import android.widget.Toast;

import java.io.File;

import go.jimsdk.*;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        Button emailVerificationButton = (Button) findViewById(R.id.emailVerificationButton);
        final EditText usernameEditText = (EditText) findViewById(R.id.usernameEditText);
        final EditText passwordEditText = (EditText) findViewById(R.id.passwordEditText);
        Button registerButton = (Button) findViewById(R.id.registerButton);
        Button loginButton = (Button) findViewById(R.id.loginButton);
        final TextView useridTextView = (TextView) findViewById(R.id.useridTextView);

        File cacheDir = this.getApplicationContext().getExternalCacheDir();
        File cookieFile = new File(cacheDir, "client_cookie_jar");

        Client client = null;

        try {
            client = Jimsdk.newClient("http://api2.jimyun.com", 23, "iu3TKjwRUCGfIwtTH9gXeYsq", "kJek81coyFG4V3eSg79b82HU", cookieFile.getPath());

            if (client.hasValidSession()) {
                client.sendUserInfoAsync(0, 0, new UserInfoResponseListener() {
                    @Override
                    public void onFailure(ResponseError responseError) {

                    }

                    @Override
                    public void onSuccess(UserInfoResponse userInfoResponse) {
                        if (userInfoResponse.getError() == null) {
                            useridTextView.setText(Long.toString(userInfoResponse.getID()));
                        }
                    }
                });
            }
        } catch (Exception e) {
            e.printStackTrace();
        }

        final Client finalClient = client;

        emailVerificationButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                finalClient.sendVerifyEmailAsync("yangjingtian@oudmon.com", new VerifyEmailResponseListener() {
                    @Override
                    public void onFailure(ResponseError responseError) {
                        Toast.makeText(getApplicationContext(), "Sent verification email - Failed. " + responseError.getMessage(), Toast.LENGTH_LONG).show();
                    }

                    @Override
                    public void onSuccess(VerifyEmailResponse verifyEmailResponse) {
                        if (verifyEmailResponse.getResult()) {
                            Toast.makeText(getApplicationContext(), "Sent verification email - OK.", Toast.LENGTH_LONG).show();
                        } else {
                            Toast.makeText(getApplicationContext(), "Sent verification email - Failed.", Toast.LENGTH_LONG).show();
                        }
                    }
                });
            }
        });

        registerButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                RegisterParams registerParams = Jimsdk.newRegisterParams();
                registerParams.setUsername(usernameEditText.getText().toString());
                registerParams.setPassword(passwordEditText.getText().toString());

                finalClient.sendRegisterAsync(registerParams, new RegisterResponseListener() {
                    @Override
                    public void onFailure(ResponseError responseError) {
                        Toast.makeText(getApplicationContext(), "Register - Failed. " + responseError.getMessage(), Toast.LENGTH_LONG).show();
                    }

                    @Override
                    public void onSuccess(RegisterResponse registerResponse) {
                        if (registerResponse.getError() != null) {
                            Toast.makeText(getApplicationContext(), "Register - Failed. " + registerResponse.getError().getMessage(), Toast.LENGTH_LONG).show();
                        } else {
                            useridTextView.setText(Long.toString(registerResponse.getID()));
                            Toast.makeText(getApplicationContext(), "Register - OK.", Toast.LENGTH_LONG).show();
                        }
                    }
                });
            }
        });

        loginButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                LoginParams loginParams = Jimsdk.newLoginParams();
                loginParams.setUsername(usernameEditText.getText().toString());
                loginParams.setPassword(passwordEditText.getText().toString());

                finalClient.sendLoginAsync(loginParams, new LoginResponseListener() {
                    @Override
                    public void onFailure(ResponseError responseError) {
                        Toast.makeText(getApplicationContext(), "Login - Failed. " + responseError.getMessage(), Toast.LENGTH_LONG).show();
                    }

                    @Override
                    public void onSuccess(LoginResponse loginResponse) {
                        if (loginResponse.getError() != null) {
                            Toast.makeText(getApplicationContext(), "Login - Failed. " + loginResponse.getError().getMessage(), Toast.LENGTH_LONG).show();
                        } else {
                            useridTextView.setText(Long.toString(loginResponse.getID()));
                            Toast.makeText(getApplicationContext(), "Login - OK.", Toast.LENGTH_LONG).show();
                        }
                    }
                });
            }
        });
    }
}
