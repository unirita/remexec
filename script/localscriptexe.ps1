$remotehost = $Args[0]
$username = $Args[1]
$passwd = $Args[2]
$execution= $Args[3]

if ($Args.Length -ne 4) {
    echo "Invalid argument."
    exit 1
}

$password = ConvertTo-SecureString $passwd -asplaintext -force
$cred = New-Object System.Management.Automation.PSCredential($username, $password)
$session = New-PSSession -ComputerName $remotehost -Credential $cred

$exelist = -split $execution
$exefile = $exelist[0]
$exelist[0] = "null"
$exearglist =  @()

foreach($val in $exelist){
    if ($val -ne "null"){
        $exearglist += $val
     }
}

$stdout = invoke-Command -Session $session -FilePath $exefile -ArgumentList $exearglist 
$result = $?

Remove-PSSession -Session $session

echo $stdout

if ($result -ne $true){
    echo "ExitStatus 1"
    exit 1
}

exit 0





