//
//  ViewController.swift
//  JimSdkExample
//
//  Created by 杨景天 on 27/09/2016.
//  Copyright © 2016 JinnSky. All rights reserved.
//

import UIKit
import Jimsdk

class ViewController: UIViewController {
    
    struct SdkConstants {
        static let ClusterURL = "http://api2.jimyun.com"
        static let AppID = 23
        static let JimAppID = "iu3TKjwRUCGfIwtTH9gXeYsq"
        static let JimAppSecret = "kJek81coyFG4V3eSg79b82HU"
    }
    
    var client: GoJimsdkClient? = nil
    
    @IBOutlet weak var emailVerificationButton: UIButton!
    @IBOutlet weak var emailVerificationIndicator: UIActivityIndicatorView!
    
    @IBOutlet weak var usernameTextField: UITextField!
    @IBOutlet weak var passwordTextField: UITextField!
    @IBOutlet weak var userRegistrationButton: UIButton!
    @IBOutlet weak var statusLabel: UILabel!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        emailVerificationIndicator.isHidden = true
        
        GoJimsdkNewClient(SdkConstants.ClusterURL,
                          SdkConstants.AppID,
                          SdkConstants.JimAppID,
                          SdkConstants.JimAppSecret,
                          &client,
                          nil)
    }

    @IBAction func tapOnEmailVerficationButton(_ sender: AnyObject) {
        emailVerificationIndicator.isHidden = false
        emailVerificationIndicator.startAnimating()
        
        DispatchQueue.global().async { [weak self] in
            guard let `self` = self, let sdkClient = self.client else { return }
            
            if let responseData = sdkClient.sendVerifyEmail("yangjingtian@oudmon.com") {
                if responseData.result() {
                    DispatchQueue.main.async {
                        self.emailVerificationIndicator.stopAnimating()
                        self.emailVerificationIndicator.isHidden = true
                    }
                }
            }
        }
    }
    
    @IBAction func tapOnUserRegistrationButton(_ sender: AnyObject) {
        guard let usernameText = self.usernameTextField.text , !usernameText.isEmpty else { return }
        guard let passwordText = self.passwordTextField.text, !passwordText.isEmpty else { return }
        
        DispatchQueue.global().async { [weak self] in
            guard let `self` = self, let sdkClient = self.client else { return }
            guard let registerParams = GoJimsdkNewRegisterParams() else { return }
            
            registerParams.setUsername(usernameText)
            registerParams.setPassword(passwordText)
            
            if let responseData = sdkClient.sendRegister(registerParams) {
                if let responseError = responseData.error() {
                    DispatchQueue.main.async {
                        self.statusLabel.text = responseError.message()
                    }
                } else {
                    DispatchQueue.main.async {
                        self.statusLabel.text = String(responseData.id_())
                    }
                }
            } else {
                DispatchQueue.main.async {
                    self.statusLabel.text = "Failed"
                }
            }
        }
    }
}

