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
        Button newsFetchButton = (Button) findViewById(R.id.newsFetchButton);
        final TextView newsUrlTextView = (TextView) findViewById(R.id.newsUrlTextView);

        File cacheDir = this.getApplicationContext().getExternalCacheDir();
        File cookieFile = new File(cacheDir, "client_cookie_jar");

        Client client = null;

        try {
            client = Jimsdk.newClient("http://api2.jimyun.com", 23, "iu3TKjwRUCGfIwtTH9gXeYsq", "kJek81coyFG4V3eSg79b82HU", cookieFile.getPath());

            if (client.hasValidSession()) {
                final UserInfoResponse response = client.sendUserInfo(0, 0);

                if (response.getError() == null) {
                    useridTextView.setText(Long.toString(response.getID()));
                } else {
                    throw new Exception("Can't create client");
                }
            }
        } catch (Exception e) {
            e.printStackTrace();
        }

        final Client finalClient = client;

        emailVerificationButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new Thread(new Runnable() {
                    @Override
                    public void run() {
                        final VerifyEmailResponse response = finalClient.sendVerifyEmail("yangjingtian@oudmon.com");

                        runOnUiThread(new Runnable() {
                            @Override
                            public void run() {
                                if (response.getResult()) {
                                    Toast.makeText(getApplicationContext(), "Sent verification email - OK.", Toast.LENGTH_LONG).show();
                                } else {
                                    Toast.makeText(getApplicationContext(), "Sent verification email - Failed.", Toast.LENGTH_LONG).show();
                                }
                            }
                        });
                    }
                }).start();
            }
        });

        registerButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new Thread(new Runnable() {
                    @Override
                    public void run() {
                        RegisterParams registerParams = Jimsdk.newRegisterParams();
                        registerParams.setUsername(usernameEditText.getText().toString());
                        registerParams.setPassword(passwordEditText.getText().toString());

                        final UserInfoResponse response = finalClient.sendRegister(registerParams);

                        runOnUiThread(new Runnable() {
                            @Override
                            public void run() {
                                if (response.getError() != null) {
                                    Toast.makeText(getApplicationContext(), "Register - Failed. " + response.getError().getMessage(), Toast.LENGTH_LONG).show();
                                } else {
                                    useridTextView.setText(Long.toString(response.getID()));
                                    Toast.makeText(getApplicationContext(), "Register - OK.", Toast.LENGTH_LONG).show();
                                }
                            }
                        });
                    }
                }).start();
            }
        });

        loginButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new Thread(new Runnable() {
                    @Override
                    public void run() {
                        LoginParams loginParams = Jimsdk.newLoginParams();
                        loginParams.setUsername(usernameEditText.getText().toString());
                        loginParams.setPassword(passwordEditText.getText().toString());

                        final UserInfoResponse response = finalClient.sendLogin(loginParams);

                        runOnUiThread(new Runnable() {
                            @Override
                            public void run() {
                                if (response.getError() != null) {
                                    Toast.makeText(getApplicationContext(), "Login - Failed. " + response.getError().getMessage(), Toast.LENGTH_LONG).show();
                                } else {
                                    useridTextView.setText(Long.toString(response.getID()));
                                    Toast.makeText(getApplicationContext(), "Login - OK.", Toast.LENGTH_LONG).show();
                                }
                            }
                        });
                    }
                }).start();
            }
        });

        newsFetchButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new Thread(new Runnable() {
                    @Override
                    public void run() {
                        NewsDigestParams newsDigestParams = Jimsdk.newNewsDigestParams();
                        newsDigestParams.setFromPage(0);
                        newsDigestParams.setPageSize(5);
                        newsDigestParams.setThumbWidth(200);
                        newsDigestParams.setThumbHeight(100);
                        newsDigestParams.setLanguage("zh");

                        final NewsDigestResponse response = finalClient.sendNewsDigest(newsDigestParams);

                        runOnUiThread(new Runnable() {
                            @Override
                            public void run() {
                                if (response.getError() != null) {
                                    Toast.makeText(getApplicationContext(), "Fetching News Digest - Failed. " + response.getError().getMessage(), Toast.LENGTH_LONG).show();
                                } else {
                                    if (response.getCollection().getSize() > 0) {
                                        newsUrlTextView.setText(response.getCollection().getItemAt(0).getArticleURL());
                                    }
                                }
                            }
                        });
                    }
                }).start();
            }
        });
    }
}
